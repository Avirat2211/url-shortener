import axios from "axios";

const URL = import.meta.env.VITE_BACKEND_URL;

interface ShortenRequest {
    LongUrl: string;
    UserId: string;
}
interface ShortenResponse {
    message: string;
    short_url: string;
}
export const shortUrl = async ({ LongUrl, UserId }: ShortenRequest): Promise<ShortenResponse | null> => {
    try {
        console.log("OK");
        console.log(URL);
        const response = await axios.post<ShortenResponse>(`${URL}/create-short-url`, {
            long_url: LongUrl,
            user_id: UserId,
        });
        console.log(response)
        return response.data;
    } catch (error) {
        console.error("Error creating short URL:", error);
        return null;
    }
};