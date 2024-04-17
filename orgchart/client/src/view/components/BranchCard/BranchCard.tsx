import styles from "./BranchCard.module.css"

interface BranchCardProps {
    city: string,
    address: string,
    countOfEmployees: number
}

const BranchCard = (props: BranchCardProps) => {
    return (
        <div className={styles.branchCard}>
            <div className={styles.branchInfoWrapper}>
                <h2 className={styles.branchInfoTitle}>Подразделение</h2>
                <p>Город: {props.city}</p>
                <p>Адрес: {props.address}</p>
            </div>
        </div>
    )
}

export {BranchCard}