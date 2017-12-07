import { Injectable } from '@angular/core';
import { Post } from './post'
import {Observable} from 'rxjs/Rx'
import { catchError } from 'rxjs/operators';
import { of } from 'rxjs/observable/of'
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable()
export class PostService {
    private postsUrl = '/admin/posts/';

    constructor( private http: HttpClient ) { }


    getPosts(): Observable<Post[]> {
        return this.http.get<Post[]>(this.postsUrl);
    }

    getPost(id: number): Observable<Post> {
        const url = `${this.postsUrl}${id}`;
        return this.http.get<Post>(url).pipe(catchError(this.handleError<Post>(`getPost id={$id}`)));
    }

    putPost(data ): void {
        this.http.put('/admin/posts/'
                              , data
                             ).subscribe();
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
