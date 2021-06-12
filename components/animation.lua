local component = require('library.component')

return function(image, sheet)

    local frames = {}
    for index = 1, #sheet.frames do
        local f = sheet.frames[index]
        frames[#frames + 1] = {
            duration = f.duration,
            quad = love.graphics.newQuad(
                f.frame.x, f.frame.y,
                f.frame.w, f.frame.h,
                imageWidth, imageHeight
            )
        }
    end

    local sequences = {}
    for index = 1, #sheet.meta.frameTags do
        local t = sheet.meta.frameTags[index]
        -- Treat pingpong direction as meaning the animation should loop
        -- There's no native way to do this in Aseprite
        local shouldLoop = t.direction == 'pingpong'
        sequences[#sequences + 1] = {
            name = t.name,
            from = t.from,
            to = t.to,
            shouldLoop = shouldLoop,
            direction = shouldLoop and 'forwards' or t.direction
        }
    end

    return component('animation', {
        frames = frames,
        sequences = sequences,
        frame = 1,
        time = 0
    })
end