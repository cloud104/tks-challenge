# Desafio TKS

## Objetivo
Esse teste, nós mostrará um pouquinho como você pensa, como avalia alguns problemas simples, sabemos que "prova não prova nada", mas algumas ações nesse teste, são ações diárias em nosso time.

## O que você deve fazer nesse teste
Queremos fazer o deploy de uma aplicação simples em um cloud provider qualquer (preferimos aws ou google).
 
Tudo deve estar automatizado. Então, utilize ansible / terraform / packer.... o que achar necessário.

Seu código **deve estar versionado** e precisa conter todos os passos para que o nosso time consiga executar seu código.

## O que esperamos

- A infraestrutura deve ser imutável
- Exista um processo definido de criação de imagens
- O deploy tem que conseguir se recuperar de uma falha (seja recriando uma instância ou reiniciando o processo...)
- A execução da automação tem que ser simples e sem muitos passos manuais.
- Documentação, pode ser em inglês ou português.
- Seja possível acessar os endpoints http utilizando HTTP e HTTPS

A aplicação
=
Desenvolvemos essa aplicação simples em GO. O código está disponível aqui: https://github.com/cloud104/tks-challenge

Os binários + _hashes_ para verificação de integridade podem ser baixados aqui: https://storage.googleapis.com/tks-challenge/challenge e https://storage.googleapis.com/tks-challenge/SHA256SUMS

Como a aplicação funciona
==
A app inicia arbitrariamente na porta **8090**, "falando http" e dá listen em todas as interfaces de rede.

O endpoint `/` responde o hostname do host onde o processo é executado;

Há um endpoint `/health` **que deve ser utilizado** para fins de _healthcheck_. Esse endpoint, de forma aleatória falha por 5 minutos consecutivos (responde http 500).

