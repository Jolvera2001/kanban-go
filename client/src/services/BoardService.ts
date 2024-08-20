import {boardSchema, boardType} from "@/models/boardType";
import { IBoardService } from "@/services/iBoardService.ts";
import { ObjectId } from "bson";

export class BoardService implements IBoardService {
    private readonly baseApiUrl : string;

    constructor() {
        this.baseApiUrl = "http://localhost:8080/api/v1/"
    }

    async createBoard(boardName: string): Promise<boardType> {
        const data = { Name: boardName };

        try {
            const response = await fetch(`${this.baseApiUrl}/board`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data),
            });

            if (!response.ok) {
                throw new Error("Network response was not ok")
            }

            const json = await response.json();
            return boardSchema.parse(json);

        } catch (e) {
            console.log(e);
            return {} as boardType;
        }
    }

    async getBoards(): Promise<boardType[]> {
        try {
            const response = await fetch(`${this.baseApiUrl}/boards`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                }
            });

            if (!response.ok) {
                throw new Error("Network response was not ok")
            }

            const json = await response.json();
            return boardSchema.array().parse(json);
        } catch (e) {
            console.log(e);
            return [] as boardType[];
        }
    }

    async updateBoard(updatedBoard: boardType): Promise<string> {
        try {
            const response = await fetch(`${this.baseApiUrl}/board`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(updatedBoard)
            });

            if (!response.ok) {
                throw new Error("Network response was not ok")
            }

            return "Board Updated";
        } catch (e) {
            console.log(e);
            return "Update Failed";
        }
    }

    async deleteBoard(id: ObjectId): Promise<string> {
        const data = { ID: id };

        try {
            const response = await fetch(`${this.baseApiUrl}/board`, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            });

            if (!response.ok) {
                throw new Error("Network response was not ok")
            }

            return "Deletion Successful";
        } catch (e) {
            console.log(e);
            return "Deletion Failed";
        }
    }
}

export class BoardServiceFactory {
    public static createService(): IBoardService {
        return new BoardService();
    }
}