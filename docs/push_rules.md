<p id="top"></p>

[Voltar](/readme.md)

# Regras de Push

## 1. Faça as alterações necessárias e adicione os arquivos

Após todas as alterações no projeto é hora de enviar as alterações ao repositório remoto. Começamos adicionando os arquivos a lista de arquivos a serem enviados.

> As mudanças realizadas após a execução deste comando não serão enviadas, sendo necessário novamente o comando.

```bash
# Adicionar mudanças a lista de push

git add .
```

## 2. Faça o commit das alterações

O `Commit` cria uma espécie de encapsulamento das mudanças preparando as para o envio ao repositório remoto. É necessário passar uma mensagem com uma breve descrição das mudanças como parametro.

```bash
# Engloba as mudanças em um commit com uma mensagem

git commit -m "Mensagem de commit"
```

## 3. Faça o push das alterações para o repositório remoto

O `Push` irá enviar as alterações ao repositório remoto desde que todas as regras sejam respeitadas dentro de instantes estes arquivos estarão no repositório remoto.

```bash
# Faz o push das alterações para o repositório remoto (origin)

git push origin main  # ou o branch atual
```

## 4. Crie uma nova tag

A tag é uma maneira de controlar as versões da aplicação, sendo o primeiro número a versão principal (major), que podem causar imcompatibilidade com versões anteriores pois são realizadas diversas mudanças, melhorias e correções. O numero do meio são as versões intermediárias que incluem correções e pequenas implementações, não devem causar quebra. Já o último numero representa os `pach's` que incluem pequenas mudanças. Exemplo: v0.1.0.

```bash
# Define a tag de versionamento

git tag v0.1.0
```

## 5. Faça o push da nova tag

Assim como o push mencionado anteriormente este também envia as alterações ao repositório remoto, porém este envia a nova tag atrelada as mudanças feitas no push anterior.

```bash
# Faz o push da nova tag ao repositório remoto (origin)
git push origin v0.1.0
```

## IMPORTANTE:

Sempre antes do envio dos arquivos faça uma checagem pela documentação a fim de garantir que as informações contidas nestes arquivos contenham as informações referentes a versão que está sendo enviada.

[Voltar ao topo](#top)