export class Person {
    constructor(first, last) {
        this.first = first
        this.last = last
    }

    fullName() {
        return `${this.first} ${this.last}`
    }

    greet() {
        var element = document.querySelector('body')
        element.innerHTML = `Hallo ${this.fullName()}`
    }
}
