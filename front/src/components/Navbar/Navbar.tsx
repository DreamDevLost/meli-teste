import styles from "./Navbar.module.scss"
import Logo from "@assets/Logo_ML.png"
import {IoIosSearch} from "react-icons/io";

export default function Navbar() {
    return <div className={styles['navbar']}>
        <div className={styles['context']}>
            <img src={Logo} alt="Logo"/>
            {/*<div className={styles['inputBox']}>*/}
            <input type="text" placeholder={'Nunca dejes de buscar'}/>
            <button><IoIosSearch size={18}/></button>
            {/*</div>*/}
        </div>
    </div>
}
