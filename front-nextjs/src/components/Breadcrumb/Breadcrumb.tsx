import styles from './Breadcrumb.module.scss'

export function Breadcrumb({ links }: { links: { name: string, url?: string }[] }) {
    return (
        <ul className={styles['breadcrumb']}>
            {links.map((link, index) => (
                <li key={index}>
                    {!link.url
                        ? <span>{link.name}</span>
                        : <a href={link.url} role='link'>{link.name}</a>}
                </li>
            ))}
        </ul>
    );
}

export default Breadcrumb;
