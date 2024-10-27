
# Analisador de Temperatura Concorrente para o "The Billion Row Challenge"

Este programa em Go foi desenvolvido para enfrentar o desafio "The Billion Row Challenge", que consiste em processar eficientemente um grande volume de dados - no caso, um bilhão de linhas de medições de temperatura. O programa demonstra como a concorrência em Go pode ser utilizada para otimizar o processamento de grandes conjuntos de dados, tornando a análise rápida e eficiente.

# O Desafio "The Billion Row Challenge"

O "The Billion Row Challenge" é um desafio informal que testa a capacidade de ferramentas e técnicas de processamento de dados em lidar com grandes volumes de informação. O desafio geralmente envolve a leitura, processamento e agregação de um bilhão de linhas de dados, o que pode ser uma tarefa desafiadora para linguagens e ferramentas que não são otimizadas para lidar com essa escala. Você pode encontrar mais informações sobre o desafio no blog oficial: https://www.morling.dev/blog/one-billion-row-challenge/

# Solução com Go e Concorrência

Este programa utiliza a linguagem Go e suas primitivas de concorrência para processar as medições de temperatura de forma eficiente. As principais características da solução incluem:

* Leitura eficiente: A função carregarMedicoes lê o arquivo de dados de forma assíncrona e envia as linhas para um canal, evitando o carregamento completo do arquivo em memória.

* Processamento paralelo: A função processarMedicoes utiliza goroutines e um canal com buffer para processar as medições concorrentemente, distribuindo a carga de trabalho entre os núcleos do processador.

* Agregação segura: Um mutex (sync.Mutex) garante que a agregação dos resultados em um mapa seja feita de forma segura e consistente, mesmo com múltiplas goroutines acessando o mapa simultaneamente.

