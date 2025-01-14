/*
 * CloudLinux Software Inc 2019-2021 All Rights Reserved
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { NO_ERRORS_SCHEMA } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormBuilder } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { MessagesService } from '@services/messages.service';
import { ServicesPageService } from '@services/services-page.service';
import { MockMatDialogRef } from '@testing/mock-mat-dialog-ref';
import { MockMessageService } from '@testing/mock-messages-service';
import { MockServicesPageService } from '@testing/mock-services-page-service';

import { ServiceDatabasesService } from '@services/service-databases.service';
import { ServiceRestoresService } from '@services/service-restores.service';
import { MockServiceDatabasesService } from '@testing/mock-service-databases-service';
import { MockServiceRestoresService } from '@testing/mock-service-restores-service';
import { RestoreServiceBackupsDialogComponent } from './restore-service-backups-dialog.component';

describe('RestoreServiceBackupsDialogComponent', () => {
    let component: RestoreServiceBackupsDialogComponent;
    let fixture: ComponentFixture<RestoreServiceBackupsDialogComponent>;
    let serviceRestoresService: MockServiceRestoresService;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [RestoreServiceBackupsDialogComponent],
            providers: [
                FormBuilder,
                { provide: MessagesService, useClass: MockMessageService },
                { provide: ServicesPageService, useClass: MockServicesPageService },
                { provide: ServiceDatabasesService, useClass: MockServiceDatabasesService },
                { provide: ServiceRestoresService, useClass: MockServiceRestoresService },
                { provide: MatDialogRef, useClass: MockMatDialogRef },
                { provide: MAT_DIALOG_DATA, useValue: 'some_db' },
            ],
            schemas: [NO_ERRORS_SCHEMA]
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(RestoreServiceBackupsDialogComponent);
        component = fixture.componentInstance;
        // @ts-ignore
        serviceRestoresService = TestBed.inject(ServiceRestoresService);
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });

    describe('when form is invalid', () => {
        beforeEach(() => {
            component.formGroup.patchValue({database: ''});
            fixture.detectChanges();
        });

        it('should not restore database', () => {
            const spy = spyOn(serviceRestoresService, 'restore').and.callThrough();
            component.onSave();
            fixture.detectChanges();

            expect(spy).not.toHaveBeenCalled();
        });
    });

    describe('when form is valid', () => {
        beforeEach(() => {
            component.formGroup.patchValue({database: 'some_db'});
            fixture.detectChanges();
        });

        it('should restore database', () => {
            const spy = spyOn(serviceRestoresService, 'restore').and.callThrough();
            component.onSave();
            fixture.detectChanges();

            expect(spy).toHaveBeenCalled();
        });
    });
});
