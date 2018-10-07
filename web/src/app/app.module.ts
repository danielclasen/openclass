import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import {MatExpansionModule} from '@angular/material/expansion';
import {HttpClientModule} from "@angular/common/http";
import { ParticipantListComponent } from './participant-list/participant-list.component';
import {
  MatButtonModule,
  MatChipsModule,
  MatDividerModule, MatFormFieldModule,
  MatIconModule, MatInputModule,
  MatListModule
} from "@angular/material";
import { SessionListComponent } from './session-list/session-list.component';
import {ReactiveFormsModule} from "@angular/forms";
import {GravatarModule} from "ngx-gravatar";

@NgModule({
  declarations: [
    AppComponent,
    ParticipantListComponent,
    SessionListComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    BrowserAnimationsModule,
    ReactiveFormsModule,

    GravatarModule,

    MatExpansionModule,
    MatChipsModule,
    MatListModule,
    MatDividerModule,
    MatIconModule,
    MatButtonModule,
    MatInputModule,
    MatFormFieldModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
