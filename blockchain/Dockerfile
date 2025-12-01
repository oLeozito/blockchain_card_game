# Usa Node 18 (LTS - Super Estável)
FROM node:18-alpine

WORKDIR /app

# 1. Inicia projeto limpo
RUN npm init -y

# 2. O PULO DO GATO: Instalamos uma versão ESPECÍFICA (v2) do Hardhat
# Isso evita baixar a v3 que está quebrando sua configuração
RUN npm install --save-dev hardhat@2.22.17

# 3. Configuração Padrão CommonJS (module.exports)
# Funciona perfeitamente com Hardhat v2
RUN echo "module.exports = { solidity: '0.8.20', networks: { hardhat: { chainId: 1337 } } };" > hardhat.config.js

EXPOSE 8545

# 4. Roda o nó
CMD ["npx", "hardhat", "node", "--hostname", "0.0.0.0"]