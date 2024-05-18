import styles from './SearchProduct.module.scss'

export default function SearchProduct() {
    return (
        <div className={styles.list}>
            <div className={styles.item}>
                <img src="https://http2.mlstatic.com/D_881016-MLM51559383738_092022-I.jpg" alt="Imagem do produto"/>
                <div className={styles.info}>
                    <span className={styles.price}>
                        $ 1.000
                    </span>
                    <span className={styles.subtitle}>
                        Apple iPhone 11 Pro Max 256 Gb Verde Medianoche
                        <br/>Completo Unico!
                    </span>
                </div>
                <span className={styles.from}>
                    Capital Federal
                </span>
            </div>
            <div className={styles.item}>
                <img src="https://http2.mlstatic.com/D_881016-MLM51559383738_092022-I.jpg" alt="Imagem do produto"/>
                <div className={styles.info}>
                    <span className={styles.price}>
                        $ 1.000
                    </span>
                    <span className={styles.subtitle}>
                        Apple iPhone 11 Pro Max 256 Gb Verde Medianoche
                        <br/>Completo Unico!
                    </span>
                </div>
                <span className={styles.from}>
                    Capital Federal
                </span>
            </div>
        </div>
    );
}
