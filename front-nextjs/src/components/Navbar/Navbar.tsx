import styles from "./Navbar.module.scss";
import Image from "next/image";
import { IoIosSearch } from "react-icons/io";

export default function Navbar() {
    return <div className={styles.navbar}>
        <div className={styles.context}>
            <a href="/">
                <Image src={'/Logo_ML.png'} alt="Logo" width={53} height={36} />
            </a>
            <form action="/items">
                <input type="text" placeholder={'Nunca dejes de buscar'} name="search" />
                <button type="submit" role="button" aria-label="Buscar produtos">
                    <IoIosSearch size={18} />
                </button>
            </form>
        </div>
    </div>
}
