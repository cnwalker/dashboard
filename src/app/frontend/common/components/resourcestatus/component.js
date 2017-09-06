// Copyright 2017 The Kubernetes Dashboard Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**
 * A component that displays the chart and legend of the current status
 * (pending, failing etc.) of workload resources
 * @final
 */
export class ResourceStatusController {
  /**
   * @ngInject
   */
  constructor() {
    /** @export {Array} */
    this.resourceList;

    /** @export {boolean} */
    this.isAggregateResource;

    /** @export {string} */
    this.title;

    /** @export {Object} */
    this.resourceStats = {};

    /** @export {number} */
    this.resourceStats.success;

    /** @export {number} */
    this.resourceStats.pending;

    /** @export {number} */
    this.resourceStats.failed;

    /** @export {Array<Object>} */
    this.resourceStats.chartValues = [];

    /** @export {!Array<string>} */
    this.colorPalette = ['#00c752', '#f00', '#ff0'];
  }

  getPodStats() {
    let pods = this.resourceList.pods;

    let resourceStats = {
      'success': 0,
      'failed': 0,
      'pending': 0,
      'total': pods.length,
    };

    pods.forEach(function(pod) {
      resourceStats[pod.podStatus.status] += 1;
    });

    resourceStats.chartValues = [
      {value: resourceStats.success / resourceStats.total * 100},
      {value: resourceStats.failed / resourceStats.total * 100},
      {value: resourceStats.pending / resourceStats.total * 100},
    ];

    return resourceStats;
  }

  getAggregateResourceStats() {
    let aggregateResourceList = this.resourceList;

    let resourceStats = {'success': 0, 'failed': 0, 'pending': 0, 'total': 0}

                        aggregateResourceList.forEach(function(aggregateResource) {
                          console.log(aggregateResource);
                        });
  }

  $onInit() {
    if (this.isAggregateResource) {
      this.resourceStats = this.getAggregateResourceStats();
    } else {
      this.resourceStats = this.getPodStats();
    }
  }
}

/**
 * @type {!angular.Component}
 */
export const resourceStatusComponent = {
  templateUrl: 'common/components/resourcestatus/resourcestatus.html',
  bindings: {
    'resourceList': '<',
    'title': '<',
    'isAggregateResource': '<',
  },
  controller: ResourceStatusController,
};