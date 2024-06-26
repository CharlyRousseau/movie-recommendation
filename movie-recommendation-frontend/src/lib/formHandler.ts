import { navigate } from "svelte-routing";

const API_URL = import.meta.env.VITE_API_URL;

export async function onSubmit(url: string, body: any, successMessage: string, redirectUrl: string) {
    try {
        const response = await fetch(`${API_URL}${url}`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(body),
        });
        const data = await response.json();
        console.log(data);
        if (!response.ok) {
            return { error: data.message };
        } else {
            setTimeout(() => {
                navigate(redirectUrl);
            }, 2000);
            return {
                success: successMessage,
                token: data.token
            };
        }
    } catch (err) {
        return { error: err instanceof Error ? err.message : "An unknown error occurred" };
    }
}