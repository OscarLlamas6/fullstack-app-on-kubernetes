import { Injectable } from '@angular/core';
import { environment } from 'environments/environment';
import { HttpClient, HttpHeaders } from "@angular/common/http"
import { Router } from '@angular/router';
import { Book } from '../model/book'

@Injectable({
  providedIn: 'root'
})

export class BooksService {
  public url: string = environment.url

  constructor(private http: HttpClient, private router: Router) { }

  getOneBook(id: string) {
    return this.http.get(`${this.url}book/single/${id}`);
  }

  getAllBooks() {
    return this.http.get(`${this.url}book/all`);
  }

  addBook(book: Book) {
    return this.http.post(`${this.url}book/create`, {
      author: book.author,
      genre: book.genre,
      pages: book.pages,
      title: book.title,
      year: book.year
    });
  }

  updateBook(book: Book) {
    return this.http.put(`${this.url}book/update`,book);
  }

  deleteBook(id: string) {
    return this.http.get(`${this.url}book/delete/${id}`);
  }
}