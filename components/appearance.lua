local component = require('library.component')

return function(image, sheet)
    -- Assume sheet has 1 frame and 1 tag if no sheet is supplied
    local imageWidth = image:getWidth()
    local imageHeight = image:getHeight()
    sheet = sheet or {
        meta = {
            slices = {},
            size = {w = imageWidth, h = imageHeight},
            frameTags = {
                {name = 'default', from = 0, to = 0, direction = 'forwards'}
            }
        },
        frames = {
            {
                filename = '0',
                frame = {x = 0, y = 0, w = imageWidth, h = imageHeight},
                rotated: false,
                trimmed: false,
                spriteSourceSize: {x = 0, y = 0, w = imageWidth, h = imageHeight},
                sourceSize: {w = imageWidth, h = imageHeight},
                duration = 100
            }
        }
    }

    local c = component('appearance', {
        image = image,
        sheet = sheet
        frame = 1,
        time = 0
    })

    return c
end