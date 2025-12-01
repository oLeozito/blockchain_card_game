# Define que os comandos não geram arquivos com esses nomes
.PHONY: all up down build clean client \
        servidor1 servidor2 servidor3 \
        parar1 parar2 parar3 \
        logs1 logs2 logs3

# --- Comandos de Controle Global ---

# Sobe todo o ambiente GRADUALMENTE (Blockchain -> Líder -> Seguidores)
up:
	@echo "--- 1. Subindo a Blockchain ---"
	docker compose up --build -d blockchain
	@echo "Aguardando 5 segundos para a blockchain estabilizar..."
	@sleep 5

	@echo "--- 2. Subindo o Servidor 1 ---"
	docker compose up --build -d servidor1
	@echo "Aguardando 3 segundos..."
	@sleep 3

	@echo "--- 3. Subindo o Servidor 2 ---"
	docker compose up --build -d servidor2
	@echo "Aguardando 3 segundos..."
	@sleep 3

	@echo "--- 4. Subindo o Servidor 3 ---"
	docker compose up --build -d servidor3
	
	@echo "--- Cluster Completo e Blockchain Operante! ---"
	@echo "Use 'make logs1' para verificar o status ou 'make client' para jogar."

# Para e remove todos os contêineres
down:
	docker compose down --remove-orphans

# Apenas constrói a imagem, sem iniciar
build:
	docker compose build

# --- LIMPEZA CIRÚRGICA (SEM SUDO) ---
# Apaga apenas as pastas de dados do Raft (servidor1, servidor2, servidor3)
# PRESERVA o instrumentos.json e players.json
clean:
	@echo "Limpando dados do Raft (mantendo JSONs)..."
	docker run --rm -v "$$(pwd)/data:/data" alpine sh -c "rm -rf /data/servidor*"
	@echo "Limpeza concluída."

# --- CLIENTE (RODANDO NO DOCKER) ---
client:
	@echo "Iniciando Cliente no Docker..."
	docker run -it --rm --network host \
		-v "$$(pwd):/app" \
		-w /app \
		golang:1.24-alpine \
		go run cliente.go

# --- Comandos de Controle Individual ---

servidor1:
	docker compose up --build -d servidor1

servidor2:
	docker compose up --build -d servidor2

servidor3:
	docker compose up --build -d servidor3

parar1:
	docker compose stop servidor1
	docker compose rm -f servidor1

parar2:
	docker compose stop servidor2
	docker compose rm -f servidor2

parar3:
	docker compose stop servidor3
	docker compose rm -f servidor3

# --- Comandos de Visualização de Logs ---

logs1:
	docker compose logs -f servidor1

logs2:
	docker compose logs -f servidor2

logs3:
	docker compose logs -f servidor3