import { Component, OnInit } from '@angular/core';
import {BooksService} from "../../services/books.service";
import { filter, map } from 'rxjs/operators';
import { HttpErrorResponse } from '@angular/common/http';
import { Book, ApiResponse} from "../../model/book";

@Component({
    selector: 'app-profile',
    templateUrl: './profile.component.html',
    styleUrls: ['./profile.component.scss']
})

export class ProfileComponent implements OnInit {

    books : Book[];
    book: Book = { author:'', genre:'', pages:'', year:'', title:'', id:'' };

    constructor(private service: BooksService) { }

    ngOnInit() {
       this.getAllBooks();
    }

    getAllBooks() {
        this.service.getAllBooks().pipe(
            filter((res: ApiResponse) => {
                return res.status == 200;
            }),
            map((res) => {
                return res.data as Book[];
            })
        ).subscribe((res: Book[]) => {
            this.books = res;
        }, (res: HttpErrorResponse) => {
            console.log('error al obtener datos')
        });
    }

    saveBook() {
        this.service.addBook(this.book).subscribe((res) =>{
            this.book = { author:'', genre:'', pages:'', year:'', title:'', id:'' };
            this.getAllBooks();
        }, (res: HttpErrorResponse) => {
            console.log(res)
        });
    }
}
