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

package v1

const (
	// cluster status
	ClusterOkStatus       = "Ready"
	ClusterNotReadyStatus = "NotReady"
	ClusterFailedStatus   = "Failed"
	ClusterUnknownStatus  = "Unknown"

	// backup status
	BackupSuccessStatus = "Success"
	BackupFailedStatus  = "Failed"
	BackupUnknownStatus = "Unknown"

	// restore status
	RestoreSuccessStatus = "Success"
	RestoreFailedStatus  = "Failed"
	RestoreRunningStatus = "Running"
	RestoreUnknownStatus = "Unknown"

	Group                = "kuberlogic.com"
	apiAnnotationsGroup  = "internal." + Group
	alertEmailAnnotation = apiAnnotationsGroup + "/" + "alert-email"
)
