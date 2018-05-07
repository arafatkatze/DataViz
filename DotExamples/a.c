#include <stdio.h>
#include <gr.h>

int main(void) {
        double x[] = {0, 0.2, 0.4, 0.6, 0.8, 1.0};
        double y[] = {0.3, 0.5, 0.4, 0.2, 0.6, 0.7};
        gr_polyline(6, x, y);
        gr_axes(gr_tick(0, 1), gr_tick(0, 1), 0, 0, 1, 1, -0.01);
        // Press any key to exit
        getc(stdin);
        return 0;
}
