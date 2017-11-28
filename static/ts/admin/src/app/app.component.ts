import { Component, OnInit } from '@angular/core';
import { Post } from './post';
import {PostService} from './post.service';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
    constructor(private postService: PostService){ }
    ngOnInit(): void {
        this.getPosts();
    }
    getPosts(): void {
        this.postService.getPosts()
            .subscribe(posts => this.posts = posts)
    }
    posts : Post[];
    title = 'app';
}
