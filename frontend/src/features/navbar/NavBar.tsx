import type { JSX } from "react"
import styles from "./NavBar.module.css"
import logo from "/src/yatslogo.svg"

export const NavBar = (): JSX.Element => {
    return (
        <div>
            <div id="logo" className={styles.Logo}>
                <img src={logo} alt="logo" />
            </div>
            <ul className={styles.NavBar}>
                <li>
                    <a
                        href=""
                    >
                        About
                    </a>
                </li>
                <li>
                    <a
                        href=""
                    >
                        Services
                    </a>
                </li>
                <li>
                    <a
                        href=""
                    >
                        Contributors
                    </a>
                </li>
                <li>
                    <button className={styles.YatsButton}>sign in</button>
                </li>
            </ul>
        </div>
    )
}
