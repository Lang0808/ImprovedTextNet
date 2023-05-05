require('dotenv').config({ debug: true })

function Home(){
    console.log(process.env)
    return (
        <div>
            Home. Process env
        </div>
    )
}

export default Home;