import styles from './Breadcrumb.module.scss'

export function Breadcrumb({links}: { links: { name: string, url?: string }[] }) {
    return (
        <ul className={styles['Breadcrumb']}>
            {links.map((link, index) => (
                <li key={index}>
                    {!link.url
                        ? <span>{link.name}</span>
                        : <a href={link.url}>{link.name}</a>}
                </li>
            ))}
        </ul>
    );
}

export default Breadcrumb;
