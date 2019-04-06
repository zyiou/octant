import { Location } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { Injectable, NgModule, NgZone } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { Router, RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { ClarityModule } from '@clr/angular';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { NamespaceComponent } from './components/namespace/namespace.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { OverviewModule } from './modules/overview/overview.module';
import { InputFilterComponent } from './components/input-filter/input-filter.component';
import { NotifierComponent } from './components/notifier/notifier.component';
import { NavigationComponent } from './components/navigation/navigation.component';


@Injectable()
export class UnstripTrailingSlashLocation extends Location {
  public static stripTrailingSlash(url: string): string {
    return url;
  }
}

@NgModule({
  declarations: [
    AppComponent,
    NamespaceComponent,
    PageNotFoundComponent,
    InputFilterComponent,
    NotifierComponent,
    NavigationComponent,
  ],
  imports: [
    BrowserModule,
    ClarityModule,
    BrowserAnimationsModule,
    HttpClientModule,
    RouterModule,
    FormsModule,
    AppRoutingModule,
    OverviewModule,
  ],
  providers: [
    {
      provide: Location,
      useClass: UnstripTrailingSlashLocation
    }
  ],
  bootstrap: [
    AppComponent
  ],
})
export class AppModule {
  constructor(private ngZone: NgZone, private router: Router) {}

  navigate(commands: any[]): void {
    this.ngZone.run(() => this.router.navigate(commands)).then();
  }
}