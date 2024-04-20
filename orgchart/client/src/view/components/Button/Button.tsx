import styles from "./Button.module.css"

export enum ButtonTypes {
    Icon="Icon",
    IconText="IconText",
    Text="Text"
}
interface ButtonProps {
    id?: string,
    onClick?: () => void,
    data: string
}

const Button = (props: ButtonProps) => {
}

export {Button}