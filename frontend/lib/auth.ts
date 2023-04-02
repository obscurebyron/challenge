const IsAuth = () => {
    if (typeof window !== "undefined") {
        const cookieValue = window.document.cookie
        .split("; ")
        .find((row) => row.startsWith("auth="))
        ?.split("=")[1];

        return cookieValue === 'valid'
    }
    return false
}

export default IsAuth