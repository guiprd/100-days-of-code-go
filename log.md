# 100 Days Of Code Golang - Log

### Day 1: January 17, 2024

**Today's Progress**: 

Criei a aplicação da sequência de Fibonacci que retorna a posição na sequência que o usuário passou via terminal.

**Thoughts:** 

* Eu usei como base a estruturação de métodos vinculados a o objeto slice que conterá os elementos da sequência.
* Está implementado o método que retornará o elemento da posição requisitada pelo usuário, mas, caso ela ainda não exista, é chamado outro método que criará os novos elementos e fará a atualização  do slice.
* Não mexi em filtros que possam evitar da aplicação quebrar, como por exemplo o usuário passar valores diferentes de números inteiros maiores ou iguais a que 1

### Day 2: January 18, 2024

**Today's Progress**: 

* Criei uma UI (User Interface) básica que o usuário irá escolher o que o código fará para ele:

        1 - One element of the Fibonacci's Sequence
        2 - A selected interval of the Fibonacci's Sequence
        3 - Print the entire sequence stored
        0 - Exit application

* Refatoração de métodos do package fibonacci.
* Implementação de filtros para não quebrar a aplicação.
* Implementação de método que retorna slice da sequência total armazenada em memória.

**Thoughts:**

* Aprendi um pouco sobre Scanf e Scanln. A escolha de Scanln é a melhor, pois há interferência de input entre as ações.
* Aprendi mais sobre métodos, atribuição de valores a ponteiros
* Enfrentei problemas de import cycle tentando fazer link de pkg fibonacci para handlers e de pkg handlers para fibonacci.
* Ainda não sei se é possível testar aplicações que recebem input de usuário.
* Modifiquei o código para tentar mockar um método que faz o Scanln, mas sem sucesso. Com isso, percebi que não entendo direito o processo de mock.
* Dado a falha de mock, preciso colocar no backlog o estudo do mesmo.
<!--- 
**Link to work:** [Calculator App](http://www.example.com)


### Day 0: February 30, 2016 (Example 2)
##### (delete me or comment me out)

**Today's Progress**: Fixed CSS, worked on canvas functionality for the app.

**Thoughts**: I really struggled with CSS, but, overall, I feel like I am slowly getting better at it. Canvas is still new for me, but I managed to figure out some basic functionality.

**Link(s) to work**: [Calculator App](http://www.example.com)


### Day 1: June 27, Monday

**Today's Progress**: I've gone through many exercises on FreeCodeCamp.

**Thoughts** I've recently started coding, and it's a great feeling when I finally solve an algorithm challenge after a lot of attempts and hours spent.

**Link(s) to work**
1. [Find the Longest Word in a String](https://www.freecodecamp.com/challenges/find-the-longest-word-in-a-string)
2. [Title Case a Sentence](https://www.freecodecamp.com/challenges/title-case-a-sentence)
--->
