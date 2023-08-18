### BookTrack - Catálogo Pessoal de Leitura em Golang

O projeto BookTrack é uma aplicação desenvolvida em Golang que visa ajudar os entusiastas da leitura a registrar, organizar e analisar seus hábitos de leitura de forma eficiente e interativa. Através de uma interface intuitiva e amigável, os usuários podem manter um registro detalhado de todos os livros que já leram, além de acessar uma série de métricas e análises baseadas nesses dados.

### Recursos Principais (in progress):

- Registro de Leituras: Os usuários podem adicionar informações sobre os livros que leram, incluindo título, autor, gênero, data de início e conclusão da leitura, avaliação pessoal e notas adicionais.

- Catálogo Pessoal: O sistema organiza automaticamente os livros registrados em uma lista personalizada, permitindo que os usuários naveguem, pesquisem e filtrem facilmente por diferentes critérios, como gênero, autor ou ano de leitura.

- Métricas Detalhadas: O BookTrack gera automaticamente uma variedade de métricas e estatísticas, como o número de livros lidos por ano, média de páginas lidas, gênero mais frequente e autor favorito. Essas métricas fornecem insights valiosos sobre os padrões de leitura do usuário ao longo do tempo.

- Recomendações Personalizadas: Com base nas preferências de gênero e autores favoritos do usuário, o sistema pode oferecer sugestões de leitura, ajudando a descobrir novos livros que possam ser do interesse do usuário.

- Visualizações Gráficas: A aplicação apresenta gráficos interativos que ilustram as tendências de leitura do usuário ao longo dos anos, permitindo uma compreensão visual rápida e clara do seu progresso e preferências.

- Backup e Sincronização: Os dados de leitura são armazenados de forma segura e a aplicação oferece a opção de fazer backup dos dados localmente ou sincronizá-los com a nuvem, garantindo que as informações estejam sempre disponíveis em qualquer dispositivo.

- Interface Amigável: O BookTrack apresenta uma interface de usuário intuitiva e responsiva, tornando o processo de adicionar, editar e analisar leituras extremamente simples e agradável.


### Rodar Projeto

```
go run main.go
```

### Rodar Banco de dados

```
docker-compose up --build
```

### Arquivo JSON com os endpoints configurados

```
api-insomnia.json
```