#!/bin/bash

APP_NAME="puzzle"
LOG_FILE="monitor_puzzle.log"
SUMMARY_FILE="monitor_resumo.log"
MAX_WAIT_TIME=300  # Máximo de 5 minutos esperando
WAIT_INTERVAL=2    # Verificar a cada 2 segundos

# Função para aguardar processo aparecer
aguardar_processo() {
    local tempo_espera=0
    echo "⏳ Aguardando processo '$APP_NAME' iniciar..."
    echo "⏰ Timeout: $MAX_WAIT_TIME segundos"
    echo "🔍 Verificando a cada $WAIT_INTERVAL segundos"
    echo "----------------------------------------"
    
    while [ $tempo_espera -lt $MAX_WAIT_TIME ]; do
        PID=$(pgrep "$APP_NAME")
        
        if [ -n "$PID" ]; then
            if ps -p "$PID" > /dev/null 2>&1; then
                echo "✅ Processo encontrado! PID: $PID"
                return 0
            fi
        fi
        
        if [ $((tempo_espera % 10)) -eq 0 ]; then
            echo "⏱️  Aguardando... ${tempo_espera}s"
        fi
        
        sleep $WAIT_INTERVAL
        tempo_espera=$((tempo_espera + WAIT_INTERVAL))
    done
    
    echo "❌ Timeout: Processo '$APP_NAME' não foi encontrado em $MAX_WAIT_TIME segundos"
    return 1
}

# Função para aguardar até o processo começar a usar CPU > 0
aguardar_cpu() {
    echo "⏳ Aguardando processo começar a usar CPU..."
    while processo_existe; do
        CPU_NOW=$(ps -p "$PID" -o %cpu --no-headers | tr -d ' ' | sed 's/,/./')
        if (( $(echo "$CPU_NOW > 0" | bc -l) )); then
            echo "✅ Uso de CPU detectado: ${CPU_NOW}%"
            return 0
        fi
        sleep 0.5
    done
    echo "❌ Processo finalizou antes de usar CPU."
    return 1
}

# Função para verificar se processo ainda existe
processo_existe() {
    kill -0 "$PID" 2>/dev/null
}

# Função para gerar resumo
gerar_resumo() {
    local end_time=$(date +%s)
    local duration=$((end_time - START_TIME))
    
    local avg_cpu=0
    local avg_mem=0
    local avg_rss=0
    
    if [ $TOTAL_READINGS -gt 0 ]; then
        avg_cpu=$(echo "scale=2; $TOTAL_CPU / $TOTAL_READINGS" | bc)
        avg_mem=$(echo "scale=2; $TOTAL_MEM / $TOTAL_READINGS" | bc)
        avg_rss=$((TOTAL_RSS / TOTAL_READINGS))
    fi
    
    local hours=$((duration / 3600))
    local minutes=$(((duration % 3600) / 60))
    local seconds=$((duration % 60))
    
    echo ""
    echo "=========================================="
    echo "           RESUMO DO MONITORAMENTO"
    echo "=========================================="
    echo "Processo: $APP_NAME (PID: $PID)"
    echo "Tempo de monitoramento: ${hours}h ${minutes}m ${seconds}s"
    echo "Total de leituras: $TOTAL_READINGS"
    echo ""
    echo "--- ESTATÍSTICAS DE CPU ---"
    echo "Máximo: ${MAX_CPU}%"
    echo "Médio: ${avg_cpu}%"
    echo ""
    echo "--- ESTATÍSTICAS DE MEMÓRIA ---"
    echo "Máximo: ${MAX_MEM}%"
    echo "Médio: ${avg_mem}%"
    echo "Pico de RSS: ${MAX_RSS}MB"
    echo "Médio RSS: ${avg_rss}MB"
    echo ""
    echo "--- DETALHES ADICIONAIS ---"
    echo "Arquivo de log: $LOG_FILE"
    echo "Resumo: $SUMMARY_FILE"
    echo "Hora inicial: $(date -d @$START_TIME '+%d/%m/%Y %H:%M:%S')"
    echo "Hora final: $(date '+%d/%m/%Y %H:%M:%S')"
    echo "=========================================="
    
    echo "=== RESUMO DO MONITORAMENTO ===" >> "$SUMMARY_FILE"
    echo "Data: $(date)" >> "$SUMMARY_FILE"
    echo "Processo: $APP_NAME (PID: $PID)" >> "$SUMMARY_FILE"
    echo "Tempo de monitoramento: ${hours}h ${minutes}m ${seconds}s" >> "$SUMMARY_FILE"
    echo "Total de leituras: $TOTAL_READINGS" >> "$SUMMARY_FILE"
    echo "CPU Máximo: ${MAX_CPU}%" >> "$SUMMARY_FILE"
    echo "CPU Médio: ${avg_cpu}%" >> "$SUMMARY_FILE"
    echo "MEM Máximo: ${MAX_MEM}%" >> "$SUMMARY_FILE"
    echo "MEM Médio: ${avg_mem}%" >> "$SUMMARY_FILE"
    echo "RSS Máximo: ${MAX_RSS}MB" >> "$SUMMARY_FILE"
    echo "RSS Médio: ${avg_rss}MB" >> "$SUMMARY_FILE"
    echo "===============================" >> "$SUMMARY_FILE"
}

# Capturar sinais
trap gerar_resumo EXIT INT TERM

# Aguardar processo aparecer
if ! aguardar_processo; then
    exit 1
fi

# Aguardar até começar a consumir CPU
if ! aguardar_cpu; then
    exit 1
fi

# Variáveis para estatísticas (só agora inicializa)
START_TIME=$(date +%s)
MAX_CPU=0
MAX_MEM=0
MAX_RSS=0
TOTAL_READINGS=0
TOTAL_CPU=0
TOTAL_MEM=0
TOTAL_RSS=0
CPU_VALUES=()
MEM_VALUES=()
RSS_VALUES=()

# Cabeçalho do log
echo "=== Início do monitoramento: $(date) ===" >> "$LOG_FILE"
echo "Processo: $APP_NAME (PID: $PID)" >> "$LOG_FILE"
echo "Intervalo: 0.5 segundos" >> "$LOG_FILE"
echo "=========================================" >> "$LOG_FILE"

echo "🔍 Monitorando processo: $APP_NAME (PID: $PID)"
echo "📊 Log: $LOG_FILE"
echo "📈 Resumo: $SUMMARY_FILE"
echo "⏹️  Pressione Ctrl+C para parar"
echo "=========================================="

# Loop principal
while processo_existe; do
    TIMESTAMP=$(date '+%H:%M:%S')
    
    CPU_USAGE=$(ps -p "$PID" -o %cpu --no-headers | tr -d ' ' | sed 's/,/./')
    MEM_USAGE=$(ps -p "$PID" -o %mem --no-headers | tr -d ' ' | sed 's/,/./')
    RSS=$(ps -p "$PID" -o rss --no-headers | tr -d ' ')
    RSS_MB=$((RSS / 1024))
    
    TOTAL_READINGS=$((TOTAL_READINGS + 1))
    TOTAL_CPU=$(echo "scale=2; $TOTAL_CPU + $CPU_USAGE" | bc)
    TOTAL_MEM=$(echo "scale=2; $TOTAL_MEM + $MEM_USAGE" | bc)
    TOTAL_RSS=$((TOTAL_RSS + RSS_MB))
    
    if (( $(echo "$CPU_USAGE > $MAX_CPU" | bc -l) )); then
        MAX_CPU=$CPU_USAGE
    fi
    
    if (( $(echo "$MEM_USAGE > $MAX_MEM" | bc -l) )); then
        MAX_MEM=$MEM_USAGE
    fi
    
    if [ "$RSS_MB" -gt "$MAX_RSS" ]; then
        MAX_RSS=$RSS_MB
    fi
    
    CPU_VALUES+=("$CPU_USAGE")
    MEM_VALUES+=("$MEM_USAGE")
    RSS_VALUES+=("$RSS_MB")
    
    OUTPUT_LINE="[$TIMESTAMP] CPU: ${CPU_USAGE}% | MEM: ${MEM_USAGE}% | RSS: ${RSS_MB}MB"
    echo "$OUTPUT_LINE"
    echo "$OUTPUT_LINE" >> "$LOG_FILE"
    
    sleep 0.5
done

echo ""
echo "⚠️  Processo $APP_NAME (PID: $PID) não está mais em execução!"
gerar_resumo
