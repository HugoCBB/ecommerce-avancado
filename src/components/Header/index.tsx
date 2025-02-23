import Sliders from "../Sliders"
import "./Header.css"

const Header = () => {
    return (
        <section className="header">
            <ul>
                <a href="#">
                    <li>Home</li>
                </a>
                <a href="#">
                    <li>Produtos</li>
                </a >
                <a href="#">
                    <li>Sobre</li>
                </a>
                <a href="#">
                    <li>Login</li>
                </a>
                <a href="#">
                    <li>Cadastro</li>
                </a>
            </ul>
        <main>
            <Sliders />
        </main>
        </section>
        
    )
}

export default Header