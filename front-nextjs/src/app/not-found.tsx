import styles from './not-found.module.scss'

export default function NotFound() {
    return <div style={{ margin: 'auto auto' }}>
        <h1 className={styles.message}>404 - Page Not Found</h1>
    </div>
}