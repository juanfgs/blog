import { Component, OnInit } from '@angular/core';
import { Post } from '../post';
import {PostService} from '../post.service';
import {Observable} from 'rxjs/Rx'

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.css']
})
export class PostsComponent implements OnInit {
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
    currentPost: Post;
    onSelect(id: number): void {
        this.postService.getPost(id).subscribe(post => this.currentPost = post);
        
    }
}
