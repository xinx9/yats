//import type { JSX } from "react"
import styles from "./NavBar.module.css"

export const NavBar = () => {

    return (
        <div className={styles.NavBar}>
            <span>
                <a
                    className={styles.NavBar}
                    href=""
                >
                    Home
                </a>
                <a
                    className={styles.NavBar}
                    href=""
                >
                    About
                </a>
                <a
                    className={styles.NavBar}
                    href=""
                >
                    Contributors
                </a>
                <button className="align-right">sign in</button>
            </span>
        </div>
    )
}
