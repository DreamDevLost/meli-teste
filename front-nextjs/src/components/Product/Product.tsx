import styles from './Product.module.scss'

export default function Product({
    item
}: {
    item: {
        id: string
        title: string
        price: {
            currency: string
            amount: number
            decimals: number
        }
        picture: string
        condition: string
        free_shipping: boolean
        sold_quantity: number
        description: string
    }
}) {
    return <div className={styles.product}>
        <div className={styles.top}>
            <img src={item.picture} alt='Product' />
            <div className={styles.pricing}>
                <span className={styles.soldInfo}>{item.condition} - {item.sold_quantity} vendidos</span>
                <h1 className={styles.title}>{item.title}</h1>
                <span className={styles.price}>$ {item.price.amount.toFixed(2)}</span>
                <button>Comprar</button>
            </div>
        </div>
        <div className={styles.description}>
            <h1 className={styles.title}>Descrição do produto</h1>
            <p className={styles.text}>{item.description}</p>
        </div>
    </div>
}
