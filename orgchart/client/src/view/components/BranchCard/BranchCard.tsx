import styles from "./BranchCard.module.css"

interface BranchCardProps {
    city: string,
    address: string,
    countOfEmployees: number
}

const BranchCard = (props: BranchCardProps) => {
    return (
        <div className={styles.branchCard}>
            <div>
                <h2>Подразделение</h2>
                <p>Город: {props.city}</p>
                <p>Адресс: {props.address}</p>
            </div>
        </div>
    )
}

export {BranchCard}