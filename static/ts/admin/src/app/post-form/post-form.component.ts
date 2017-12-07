import { Component, OnInit, Input } from '@angular/core';
import { NgForm } from '@angular/forms';
import { SelectControlValueAccessor } from '@angular/forms';
import { Post } from '../post';
import { Category } from '../category';
import {PostService} from '../post.service';
import {CategoryService} from '../category.service';


@Component({
  selector: 'app-post-form',
  templateUrl: './post-form.component.html',
  styleUrls: ['./post-form.component.css']
})
export class PostFormComponent implements OnInit {

    @Input() post: Post;
    constructor(private postService: PostService, private categoryService: CategoryService) { }
    categories: Category[];
    ngOnInit() {
        this.getCategories()
    }
    
    compareCategories(c1: Category, c2: Category): boolean {
        return c1 && c2 ? c1.Id === c2.Id : c1 === c2;
    }
    onSubmit(form: NgForm) {
        this.postService.putPost(form);
    }
    getCategories(): void {

         this.categoryService.getCategories().subscribe(categories => this.categories = categories);
    }
}
