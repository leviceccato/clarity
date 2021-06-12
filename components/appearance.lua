local component = require('library.component')

return function(image)
    return component('appearance', {
        image = image
    })
end