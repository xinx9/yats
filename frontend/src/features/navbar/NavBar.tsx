import type { JSX } from "react"
import styles from "./NavBar.module.css"
import logo from "/src/yatslogo.svg"

const navItems = ["About", "Services", "Contributors"];

export const NavBar = (): JSX.Element => {
    return (
        <div>
            <div id="logo" className={styles.Logo}>
                <img src={logo} alt="logo" />
            </div>
            <ul className={styles.NavBar}>
                {navItems.map((item) => (
                    <li>
                        <a href=''>{item}</a>
                    </li>
                ))}
                <li>
                    <button className={styles.YatsButton}>sign in</button>
                </li>
            </ul>
        </div>
    )
}
