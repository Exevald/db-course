import styles from "./TextArea.module.css"

interface TextAreaProps {
    placeholder?: string
    value?: string
    id: string
}

const TextArea = ({
                      placeholder = "",
                      value,
                      id
                  }: TextAreaProps) => {
    return (
        <input type="text"
               id={id}
               placeholder={placeholder}
               className={styles.inputDefault}
               value={value}
        />
    )
}


export {TextArea}
