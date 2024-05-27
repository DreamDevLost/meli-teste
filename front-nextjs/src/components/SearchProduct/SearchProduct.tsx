import Link from 'next/link';
import styles from './SearchProduct.module.scss'

export default function SearchProduct({
    product
}: {
    product: {
        id: string,
        title: string,
        price: {
            currency: string,
            amount: number,
            decimals: number
        },
        picture: string,
    }[]
}) {
    return (
        <div className={styles.list} data-testid="product-list">
            {product.map((item, index) => (
                <div className={styles.item} key={index} data-testid="product-item">
                    <img src={item.picture} alt="Imagem do produto" />
                    <div className={styles.info}>
                        <span className={styles.price}>
                            <a href={`/items/${item.id}`}>$ {item.price.amount}</a>
                        </span>
                        <Link href={`/items/${item.id}`}>
                            <span className={styles.title} data-testid="product-title">
                                {item.title}
                            </span>
                        </Link>
                    </div>
                    <span className={styles.from}>
                        Capital Federal
                    </span>
                </div>
            ))}
        </div>
    );
}
