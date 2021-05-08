# Tutorial GoSec

Esse tutorial tem como objetivo demonstrar o uso da ferramenta Go Sec.

## O que é?

Go Sec é uma ferramenta que inspeciona o código e busca por falhas de segurança. O GoSec faz isso atraveś de uma análise
da AST (Abstract Syntax Tree) do Go.

## Passo a passo

O tutorial irá demonstrar a instalação local e o uso do Go Sec.

- Passo 1:
    Instalar o pacote do Go Sec
    ```
    go get github.com/securego/gosec/cmd/gosec
    ```

- Passo 2:
    Agora o Go Sec está disponível e já é possível realizar verificações

    Por padrão, todas as regras do Go Sec estarão habilitadas, para mais detalhes sobre as regras disponíveis, checar a documentação:

    [Uso do Go Sec](https://github.com/securego/gosec#usage)

- Passo 3:
    Para rodar o Go Sec em todos os arquivos presentes em um diretório, basta rodar o seguinte comando:

    ```
    gosec .
    ```

    Após esse comando, o Go Sec irá gerar um relatório apontando se há falhas ou não, e também irá exibir qual é a regra que foi aplicada para a detecção da(s) falha(s), caso houver.

Segue um exemplo de um relatório após rodar o Go Sec:

![Go Sec Report](gosecreport.png)

Nesse caso, o Go Sec encontrou uma vulnerabilidade de segurança no código.

É possível ver na imagem a regra violada (G404), a descrição da regra, o nível de confiança do Go Sec e o nível de severidade da vulnerabilidade.

Além disso, o Go Sec também sugere uma alteração para corrigir a vulnerabilidade.

Como foi demonstrado, o Go Sec pode ser um bom aliado para analisar o código e buscar vulnerabilidades de segurança em projetos escritos em Go.

## Tabela de regras

| Regra | Descrição |
| ----- | --------  |
| G101      | Procura por credenciais no código |
| G102      | Bind para todas as interfaces          |
| G103      | Audita o uso do bloco unsafe          |
| G104      | Audita erros não checados          |
| G106      | Audita o uso de ssh.InsecureIgnoreHostKey          |
| G107      | Url fornecida para o request HTTP como entrada "manchada"          |
| G108      | Endpoint de profiling automaticamente exposto em /debug/pprof          |
| G109      | Potencial overflow de inteiro criado por strconv.Atoi na conversão para int16/32          |
| G110      | Potencial vulnerabilidade de DoS via bomba de descompressão          |
| G201      | Construção de query SQL usando formatação de string          |
| G202      | Construção de query SQL usando concatenação de string          |
| G203      | Uso de dados sem escape em template HTML          |
| G204      | Audita o uso de execução de comando          |
| G301      | Permissões ruins de arquivo ao criar um diretório          |
| G302      | Permissões ruins de arquivo usado com chmod          |
| G303      | Criando arquivo temporário usando um caminho previsível          |
| G304      | Caminho de arquivo fornecido como entrada "manchada"          |
| G305      | Passagem de arquivo ao extrair arquivo tar/zip          |
| G306      | Permissões ruins de arquivo quando escreve para um novo arquivo          |
| G307      | Defere um método que retorna um erro         |
| G401      | Detecta o uso de DES, RC4, MD5 ou SHA1          |
| G402      | Procura por configurações ruins de conexões TLS          |
| G403      | Garante o comprimento mínimo de 2048 bits para a chave RSA          |
| G404      | Fonte insegura de número aleatório (rand)          |
| G501      | Lista de import bloqueado: crypto/md5         |
| G502      | Lista de import bloqueado: crypto/des          |
| G503      | Lista de import bloqueado: crypto/rc4         |
| G504      | Lista de import bloqueado: net/http/cgi         |
| G505      | Lista de import bloqueado: crypto/sha1 |
| G601      | Aliasing de memória implícito dos itens em um uma declaração de range          |


## Licença

[![CC BY-SA 4.0][cc-by-sa-shield]][cc-by-sa]

Esse trabalho está licenciado sob a licença
[Creative Commons Attribution-ShareAlike 4.0 International License][cc-by-sa].

[![CC BY-SA 4.0][cc-by-sa-image]][cc-by-sa]

[cc-by-sa]: http://creativecommons.org/licenses/by-sa/4.0/
[cc-by-sa-image]: https://licensebuttons.net/l/by-sa/4.0/88x31.png
[cc-by-sa-shield]: https://img.shields.io/badge/License-CC%20BY--SA%204.0-lightgrey.svg
    
