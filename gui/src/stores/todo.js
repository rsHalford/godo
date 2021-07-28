import { writable } from 'svelte/store';

export const todo = writable({
  ID: "0",
  priority: false,
  status: false,
  title: "Name your todo",
  body: `# Use Markdown Syntax to make your Todos easier to read!

# Heading 1

## Heading 2

### Heading 3

#### Heading 4

##### Heading 5

###### Heading 6

**This is a bold text using double asterisks**

__This is a bold text using double underscore__

*This is italic text using single asterisk*

_This is italic text using single underscore_

**_This is bold and italic with a combination_**

> How do you like my quote?
> 
> Does it inspire you?

1. This is an ordered list
2. Continued

\`function parseMarkdown(markdownText)\`

\`\`\`
// This should now be a codeblock
func getMarkedDown() string {
	s := "You got marked down!"
	fmt.Printf("%s", s)
	message := s
	return message
}
\`\`\`

![This is an image alt text](https://i.imgur.com/KsCNBdf.jpg)

[This is a link to the GoDo about page](https://todo.xhalford.com)`,
});
