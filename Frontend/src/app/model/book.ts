export interface Book {
    id: string;
    author: string;
    genre: string;
    pages: string;
    title: string;
    year: string;
}

export interface ApiResponse {
    status: number,
    message: string,
    data: any,
}