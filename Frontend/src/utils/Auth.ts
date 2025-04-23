import Cookies from 'js-cookie';


export const setToken= async (token: string) => {
    Cookies.set("token", token);
}

export const getToken = () => {
    return Cookies.get("token");
}

export const removeToken = () => {
    Cookies.remove("token");
}

const getpayload = (token: string) => {
    const payload = token ? token.split(".")[1] : "";
    return JSON.parse(atob(payload));
}

export const getUserId = () => {    
    const token = getToken();
    if (token) {
        const payload = getpayload(token);
        return payload.user_id;
    }
    return null;
}

export const getUserRole = () => {
    const token = getToken();
    if (token) {
        const payload = getpayload(token);
        return payload.role;
    }
    return null;
}