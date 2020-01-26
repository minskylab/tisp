import { html, define, Hybrids } from "hybrids";

interface SimpleCounter extends HTMLElement {
  count: number;
}

export function increaseCount(host: SimpleCounter) {
  host.count += 1;
}

export const SimpleCounter: Hybrids<SimpleCounter> = {
  count: 0,
  render: ({ count }) => html`
    <span>Count: ${count}</span>
    <br />
    <button onclick="${increaseCount}">Add +1</button>
  `
};

define("simple-counter", SimpleCounter);
