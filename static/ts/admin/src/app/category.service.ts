import { Injectable } from '@angular/core';
import { Category } from './category'
import {Observable} from 'rxjs/Rx'
import { catchError } from 'rxjs/operators';
import { of } from 'rxjs/observable/of'
import { HttpClient, HttpHeaders } from '@angular/common/http';


@Injectable()
export class CategoryService {
    private categoriesUrl = '/admin/categories/';

    constructor( private http: HttpClient ) { }

    getCategories(): Observable<Category[]> {
        return this.http.get<Category[]>(this.categoriesUrl);
    }
    getCategory(id: number): Observable<Category> {
        const url = `${this.categoriesUrl}${id}`;
        return this.http.get<Category>(url).pipe(catchError(this.handleError<Category>(`getCategory id={$id}`)));
    }
    private handleError<T> (operation = 'operation', result?: T) {
      return (error: any): Observable<T> => {
        // TODO: send the error to remote logging infrastructure
        console.error(error); // log to console instead
        // Let the app keep running by returning an empty result.
        return of(result as T);
      };
    }

}
