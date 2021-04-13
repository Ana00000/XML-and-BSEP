var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    }
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
import Vue from 'vue';
import { EJComponentDecorator } from '@syncfusion/ej2-vue-base';
var ResourcesDirective = /** @class */ (function (_super) {
    __extends(ResourcesDirective, _super);
    function ResourcesDirective() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    ResourcesDirective.prototype.render = function () {
        return;
    };
    ResourcesDirective = __decorate([
        EJComponentDecorator({})
    ], ResourcesDirective);
    return ResourcesDirective;
}(Vue));
export { ResourcesDirective };
export var ResourcesPlugin = {
    name: 'e-resources',
    install: function (Vue) {
        Vue.component(ResourcesPlugin.name, ResourcesDirective);
    }
};
/**
 * `e-resources` directive represent a resources of the VueJS Schedule.
 * It must be contained in a Schedule component(`ejs-schedule`).
 * ```vue
 * <ejs-schedule>
 *   <e-resources>
 *    <e-resource field='RoomId' name='Rooms'></e-resource>
 *    <e-resource field='OwnerId' name='Owners'></e-resource>
 *   </e-resources>
 * </ejs-schedule>
 * ```
 */
var ResourceDirective = /** @class */ (function (_super) {
    __extends(ResourceDirective, _super);
    function ResourceDirective() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    ResourceDirective.prototype.render = function () {
        return;
    };
    ResourceDirective = __decorate([
        EJComponentDecorator({})
    ], ResourceDirective);
    return ResourceDirective;
}(Vue));
export { ResourceDirective };
export var ResourcePlugin = {
    name: 'e-resource',
    install: function (Vue) {
        Vue.component(ResourcePlugin.name, ResourceDirective);
    }
};
