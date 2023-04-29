# ImprovedTextNet

1. Motivation
- This is an improved version of a project I used for probation in Zalo, VNG.
- This project doesnt contain any piece of code in the project in VNG. 
- I do this project because I want to see how better I am after 1 year in VNG

2. Functionalities <br/>
2.1. Overall functionalities
- This project is a simple social network. You can make friends with other users, create post and share your ideas with them, and you can see what your friend are thinking about and interact with them.

3. Architectures
- This project is designed in microservice architecture
- My initial design:
![image desc](docs/InitialArchitecture.jpg)
3.1. Front end
- In my project in VNG, I used server side rendering, and the front end is only native HTML, CSS, JavaScript
- Now, I will use ReactJS for front end
- The reason why I change to ReactJS is performance. ReactJS has a good performance because it manipiulate VirtualDOM, not real DOM, so it gains better performance if front end change a little bit at a time. ReactJS only change modified parts in page, not change a whole page.
3.2. Back end
3.2.1. UserService
- This service is responsible for managing user's information, including Id, UserName, Avatar, Password
- User's password will be hashed/encrypted
- In VNG, my project uses NoSQL to store user's information. Now, I use MySQL to do it. This change occurs because I want to utilize the unique constraints of MySQL.
- Initial design of UserService:
![image desc](docs/InitialDesignUserService.jpg)
