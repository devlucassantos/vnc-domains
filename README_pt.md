# vnc-domains

🌍 *[English](README.md) ∙ [Português](README_pt.md)*

`vnc-domains` é o serviço responsável por centralizar os domínios e as regras de negócio da plataforma
[Você na Câmara (VNC)](#você-na-câmara). Neste repositório, você encontrará o código-fonte desses domínios e regras de 
negócio, utilizando como principal tecnologia Go. Deste modo, esse repositório serve apenas como um serviço auxiliar
para os serviços [`vnc-summarizer`](https://github.com/devlucassantos/vnc-summarizer) e
[`vnc-api`](https://github.com/devlucassantos/vnc-api).

## Você na Câmara

Você na Câmara (VNC) é uma plataforma de notícias desenvolvida para simplificar e tornar acessíveis às proposições
legislativas que tramitam na Câmara dos Deputados do Brasil. Por meio do uso de Inteligência Artificial, a plataforma
sintetiza o conteúdo desses documentos legislativos, transformando informações técnicas e complexas em resumos objetivos
e claros para a população em geral.

Este projeto integra o Trabalho de Conclusão de Curso dos desenvolvedores da plataforma e foi concebido com base
em arquiteturas como a hexagonal e a de microsserviços. A solução foi organizada em diversos repositórios, cada um com
responsabilidades específicas dentro do sistema:

* [`vnc-databases`](https://github.com/devlucassantos/vnc-databases): Responsável por gerenciar a infraestrutura de
  dados da plataforma. Principais tecnologias utilizadas: PostgreSQL, Redis, Liquibase e Docker.
* [`vnc-pdf-content-extractor-api`](https://github.com/devlucassantos/vnc-pdf-content-extractor-api): Responsável por
  realizar a extração de conteúdo dos PDFs utilizados pela plataforma. Principais tecnologias utilizadas: Python,
  FastAPI e Docker.
* [`vnc-domains`](https://github.com/devlucassantos/vnc-domains): Responsável por centralizar os domínios e regras de
  negócio da plataforma. Principal tecnologia utilizada: Go.
* [`vnc-summarizer`](https://github.com/devlucassantos/vnc-summarizer): Responsável pelo software que extrai os dados e
  sumariza as proposições disponibilizadas na plataforma. Principais tecnologias utilizadas: Go, PostgreSQL, Amazon Web
  Services (AWS) e Docker.
* [`vnc-api`](https://github.com/devlucassantos/vnc-api): Responsável por disponibilizar os dados para o frontend da
  plataforma. Principais tecnologias utilizadas: Go, Echo, PostgreSQL, Redis e Docker.
* [`vnc-web-ui`](https://github.com/devlucassantos/vnc-web-ui): Responsável por fornecer a interface web da plataforma.
  Principais tecnologias utilizadas: TypeScript, SCSS, React, Vite e Docker.
