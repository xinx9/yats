import type { JSX } from "react"
import styles from "./NavBar.module.css"
import logo from "/src/yatslogo.svg"

const navItems = ["About", "Services", "Contributors"];
const navAnchors = ["#AboutYats", "#Services", "#Contributors"];

export const NavBar = (): JSX.Element => {
    return (
        <div className={styles.FixedTop}>
            <div id="logo" className={styles.Logo}>
                <img src={logo} alt="logo" />
            </div>
            <ul className={styles.NavBar}>
                {navItems.map((item, index) => (
                    <li>
                        <a href={navAnchors[index]}>{item}</a>
                    </li>
                ))}
                <li>
                    <button className={styles.YatsButton}>sign in</button>
                </li>
            </ul>
        </div>
    )
}
