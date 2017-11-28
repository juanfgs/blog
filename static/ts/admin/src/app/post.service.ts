import { Injectable } from '@angular/core';
import { Post } from './post'
import {Observable} from 'rxjs/Rx'
import { of } from 'rxjs/observable/of'
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable()
export class PostService {
    private postsUrl = '/admin/posts/';

    constructor( private http: HttpClient ) { }


    getPosts(): Observable<Post[]> {
        return this.http.get<Post[]>(this.postsUrl);
    }
}
