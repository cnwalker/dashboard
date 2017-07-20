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

import {stateName} from 'search/state';

/**
 * @final
 */
export class SearchController {
  /**
   * @param {!ui.router.$state} $state
   * @ngInject
   */
  constructor($state) {
    /** @private {!ui.router.$state} */
    this.state_ = $state;
    this.query = '';

    // Live search variables
    this.last_search_time = 0;
    this.search_interval = 3000;  // (In milliseconds)
    this.queuedQuery;
  }

  /**
   * @export
   */
  queueThrottledSubmission() {
    var current_time = Math.round(new Date());
    var current_interval = current_time - this.last_search_time;

    this.last_search_time = current_time;

    if (current_interval >= this.search_interval) {
      this.state_.go(stateName, {q: this.query});
    } else {
      if (!this.queuedQuery) {
        this.queuedQuery = setTimeout(function() {
          this.state_.go(stateName, {q: this.query});
          this.queuedQuery = null;
        }.bind(this), Math.min(current_interval, this.search_interval));
      }
    }
  }

  /**
   * @export
   */
  submit() {
    clearTimeout(this.queuedQuery);
    this.state_.go(stateName, {q: this.query});
  }
}

/** @type {!angular.Component} */
export const searchComponent = {
  controller: SearchController,
  templateUrl: 'chrome/search/search.html',
};
