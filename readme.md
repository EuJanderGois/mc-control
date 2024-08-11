# MC Control - v0.0.1-prototype

> Este nome é alernativo e provisório.

> Sorry if there are still parts of the documentation without translation.

## O que é

Esta é uma versão de protótipo de uma ferramenta que propoe solucionar problemas relacionados ao desenvolvimento de servidores minecraft por pessoas de fora do ramo de desenvolvimento de software ou programação.

Vamos imaginar um ambiente de uma empresa que gerencia um servidor. Nesta equipe existem duas pessoas que trabalham de forma remota e são responsáveis por verificar a jogabilidade e realizar pequenas mudanças no servidor.

Neste cenário existem dois caminhos prováveis. Cada um deles executa o servidor em sua propria maquina de forma que copiem os parametros do servidor original e depois implementam estas mudanças no servidor de produção.

Esta configuração de servidores em maquina local, para testes ou outras finalidades, trazem diversos desafios, como pesquisar no youtube porque o Java não é reconhecido, ou o que é Java.

Desafios que programadores enfrentam frequentemente, portanto já não podem mais ser considerados desafios. Porém ainda existe uma dificuldade ali e esta aplicação se propõe a solucionar isto.

## Como funciona

MC Control é construido em [GoLang]() e cria uma ponte entre o usuário e a configuração do servidor.

De forma mais detalhada, MCC cria um servidor intermediário que é servido em `Local Host` na porta `3333`. A intenção é criar uma interface amigavel desde a instalação, o que deve ser implementado futuramente.

Por enquanto a instalação inicial do servidor [Spigot]() ainda precisa ser manual. Já na aplicação MCC o diretório do servidor [Spigot]() deve ser indicado para o funcionamento.

Ao executar `start.exe`, uma aba do navegador deve ser aberta em conjunto com o terminal do servidor.

No momento a única dependencia é o arquivo `start.exe` e os arquivos estáticos que devem estar armazenados em `services/cmd/web`. Onde devem estar arquivos HTML, CSS, etc.

Na janela aberta no navegador ou no endereço http://localhost:3333/ as configurações do servidor podem ser visualizadas e modificadas em uma interface amigável que tem sua inspiração em interfaces como Aternos e Pterodactyl.
