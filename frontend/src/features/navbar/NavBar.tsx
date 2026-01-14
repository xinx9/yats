import type { JSX } from "react"
import "./NavBar.scss"
import logo from "/src/yatslogo.svg"

const navItems = ["About", "Services", "Contributors"];
const navAnchors = ["#AboutYats", "#Services", "#Contributors"];

function logoOnClick() {
    //alert("clicked");
}

export const NavBar = (): JSX.Element => {
    return (
        <div className={"FixedTop"}>
            <div id="logo" className={"LogoContainer"}>
                <button className={"elevated"} onClick={logoOnClick}><img src={logo} alt="logo" /></button>
            </div>
            <ul className={"NavBar"}>
                {navItems.map((item, index) => (
                    <li>
                        <a href={navAnchors[index]}>{item}</a>
                    </li>
                ))}
                <li>
                    <button className={"ActionButton"}>sign in</button>
                </li>
            </ul>
        </div>
    )
}
