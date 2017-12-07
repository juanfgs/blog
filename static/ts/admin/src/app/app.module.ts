import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { AppComponent } from './app.component';
import { PostService } from './post.service';
import { CategoryService } from './category.service';
import { PostsComponent } from './posts/posts.component';

import { HttpClientModule  } from '@angular/common/http';
import { PostFormComponent } from './post-form/post-form.component';

@NgModule({
  declarations: [
    AppComponent,
    PostsComponent,
    PostFormComponent
  ],
  imports: [
      BrowserModule,
      FormsModule,
      HttpClientModule
  ],
    providers: [PostService,CategoryService],
  bootstrap: [AppComponent]
})
export class AppModule { }
